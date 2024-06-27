package mps


import (
  "testing"
)


func TestSearch(t *testing.T) {
  dictionary := Dictionary{"word1": "meaning1", "word2": "meaning2"}

  t.Run("known word", func(t *testing.T){
    got, _ := dictionary.Search("word2")
    want := "meaning2"
    assertStrings(t, got, want)
  })

  t.Run("unknown word", func(t *testing.T){
    _, err := dictionary.Search("wordX")
    if err == nil {
      t.Fatal("expected to get an error. didn't get")
    }
    assertError(t, err, ErrNotFound)
  })
}


func TestAdd(t *testing.T) {
  t.Run("new word", func(t *testing.T){
    dict := Dictionary{}
    err := dict.Add("word3", "meaning3")
    assertError(t, err, nil)
    assertDefinition(t, dict, "word3", "meaning3")
  })

  t.Run("existing word", func(t *testing.T){
    word := "repeat"
    meaning := "again"
    dict := Dictionary{word: meaning}
    err := dict.Add(word, "some other meaning")
    if err == nil {
      t.Fatal("expected to get an error. didn't get")
    }
    assertError(t, err, ErrWordExists)
    assertDefinition(t, dict, word, "again")
  })
}


func TestUpdate(t *testing.T) {
  t.Run("existing word", func(t *testing.T){
    word := "word"
    meaning := "meaning"
    dict := Dictionary{word: meaning}
    newMeaning := "newMeaning"
    err := dict.Update(word, newMeaning)
    assertError(t, err, nil)
    assertDefinition(t, dict, word, newMeaning)
  })

  t.Run("new word", func(t *testing.T){
    word := "word"
    meaning := "meaning"
    dict := Dictionary{}
    err := dict.Update(word, meaning)
    assertError(t, err, ErrWordDoesntExist)
  })
}

func TestDelete(t *testing.T) {
  word := "word"
  meaning := "meaning"
  dict := Dictionary{word: meaning}
  dict.Delete(word)
  _, err := dict.Search(word)
  assertError(t, err, ErrNotFound)
}


func assertDefinition(t testing.TB, dict Dictionary, key, value string) {
  t.Helper()
  got, err := dict.Search(key)
  if err != nil {
    t.Fatal("blah")
  }
  assertStrings(t, got, value)
}


func assertError(t testing.TB, got, want error) {
  t.Helper()
  if got != want {
    t.Errorf("got error: %q, want error: %q", got, want)
  }
}


func assertStrings(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got: %q, want: %q", got, want)
  }
}
