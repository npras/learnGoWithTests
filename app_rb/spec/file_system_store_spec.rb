require './spec/spec_helper.rb'
require './db/file_system_store.rb'

describe Db::FileSystemStore do

  before do
    @store = nil
    Tempfile.create('db-test') do |f|
      f.write content
      @store = Db::FileSystemStore.new f
    end
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
    Tempfile.create('db-test-empty') do |f|
      f.write ''
      @store = Db::FileSystemStore.new f
    end
    name = 'NewLeo'
    assert_nil @store.get_player_score(name)
    @store.record_win(name)
    assert_equal 1, @store.get_player_score(name)
  end

end

def content
  <<~EOF
      [{"Name": "Joe", "Wins": 1},
      {"Name": "Cleo", "Wins": 10},
      {"Name": "Chris", "Wins": 32}]
  EOF
end
