require './spec/spec_helper.rb'
require './lib/controllers/player_controller.rb'

describe PlayerController do

  include Rack::Test::Methods

  def app = PlayerController
  def store = app.settings.store


  describe "integration" do
    it "works together" do
      Tempfile.create('db-test') do |f|
        f.write '[]'
        app.set :store, Db::FileSystemStore.new(f)
      end
      #app.set :store, Db::InMemoryStore.new

      name = 'pepperx'
      assert_nil store.get_player_score(name)

      # GET player(score)
      get "/players/#{name}"
      assert last_response.not_found?

      # POST player(win)
      7.times { post "/players/#{name}" }

      # GET player(score) again
      get "/players/#{name}"
      assert last_response.ok?
      assert_equal '7', last_response.body

      # GET league
      get '/league'

      assert last_response.ok?
      assert_equal 'application/json', last_response.content_type
      got = JSON.parse last_response.body
      assert_equal [[name, 7]], got
    end
  end

end
