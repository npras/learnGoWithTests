require './spec/spec_helper.rb'
require './db/file_system_store.rb'

def content
  <<~EOF
      [{"Name": "Joe", "Wins": 1},
      {"Name": "Cleo", "Wins": 10},
      {"Name": "Chris", "Wins": 32}]
  EOF
end

describe Db::FileSystemStore do

  before do
    @f = Tempfile.new('db-test').tap { _1.write content }
    @store = Db::FileSystemStore.new @f
  end

  after do
    @f.close
    @f.unlink
  end

  it '#get_league sorted' do
    want = [[ "Chris", 32 ],
            [ "Cleo", 10 ],
            [ "Joe", 1 ]]
    assert_equal want, @store.get_league
  end

  it '#get_player_score' do
    assert_equal 32, @store.get_player_score('Chris')
    assert_nil @store.get_player_score('Non-Existent')
  end

  describe '#record_win' do
    it 'records for existing player' do
      name = 'Chris'
      8.times { @store.record_win(name) }
      assert_equal 40, @store.get_player_score(name)
    end

    it 'records for new player' do
      name = 'NewLeo'
      @store.record_win(name)
      assert_equal 1, @store.get_player_score(name)
    end
  end

  it "works with empty file" do
    @f = Tempfile.new('db-test-empty').tap { _1.write '' }
    store = Db::FileSystemStore.new @f
    name = 'NewLeo'
    assert_nil store.get_player_score(name)
    store.record_win(name)
    assert_equal 1, store.get_player_score(name)
  end

end
