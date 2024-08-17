ENV['APP_ENV'] = 'test'

require './player_server'
require 'minitest/autorun'
require 'rack/test'

describe PlayerServer do

  include Rack::Test::Methods

  def app = PlayerServer
  def store = app.settings.store

  describe "GET player_score" do
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

  describe "POST record_win" do
    it "records wins when POSTed" do
      name = 'pepperx'
      assert_nil store.get_player_score(name)
      7.times { post "/players/#{name}" }
      assert_equal 201, last_response.status
      assert_equal 7, store.get_player_score(name)
    end
  end

end
