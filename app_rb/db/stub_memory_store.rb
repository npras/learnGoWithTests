module Db
  class StubMemoryStore

    attr_accessor :h, :win_calls

    def initialize(h = {})
      @h = h
      @win_calls = []
    end

    def get_player_score(name) = h[name]
    def record_win(name) = win_calls << name
    def get_league = h
  end
end
