ENV['APP_ENV'] = 'test'

require 'minitest/autorun'
require 'rack/test'
require './player_server.rb'

describe PlayerServer do

  include Rack::Test::Methods

  def app = PlayerServer
  def store = app.settings.store


  describe "GET_player_score" do
    before do
      h = {
        "pepper" => 20,
        "kyle" => 27,
      }
      app.set :store, Db::InMemoryStore.new(h)
    end

    it "gets pepper's score" do
      get '/players/pepper'
      assert_equal '20', last_response.body
      assert last_response.ok?
    end

    it "gets kyle's score" do
      get '/players/kyle'
      assert_equal '27', last_response.body
      assert last_response.ok?
    end

    it "return's 404 for notfound player" do
      get '/players/notfound'
      assert_equal 404, last_response.status
    end
  end


  describe "POST_record_win" do
    before do
      app.set :store, Db::StubMemoryStore.new
    end

    it "calls record_win each time" do
      name = 'pepperx'
      assert_nil store.get_player_score(name)
      assert_equal 0, store.win_calls.size
      7.times { post "/players/#{name}#{_1+1}" }
      assert_equal 7, store.win_calls.size
      assert_equal name+'1', store.win_calls.first
      assert_equal name+'7', store.win_calls.last
    end
  end


  describe "GET_league" do
    it "returns league table as JSON" do
      want = {
        "Cleo" => 32,
        "Chris" => 20,
        "Tiest" => 14,
      }
      app.set :store, Db::StubMemoryStore.new(want)
      get '/league'
      assert last_response.ok?
      assert_equal 'application/json', last_response.headers['content-type']
      got = JSON.parse last_response.body
      assert_equal want, got
    end
  end

end


module Db
  class StubMemoryStore

    attr_accessor :h, :win_calls

    def initialize(h = {})
      @h = h
      @win_calls = []
    end

    def get_player_score(name) = h[name]
    def record_win(name) = win_calls << name

  end
end
