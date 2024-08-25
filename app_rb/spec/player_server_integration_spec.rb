require './spec/spec_helper.rb'
require './lib/controllers/player_controller.rb'

describe PlayerController do

  include Rack::Test::Methods

  def app = PlayerController
  def store = app.settings.store


  describe 'integration' do
    it 'works together' do
      Tempfile.create('db-test') do |f|
        f.write '[]'
        app.set :store, Db::FileSystemStore.new(f)
      end

      name1 = 'pepperx'
      name2 = 'kylex'
      name3 = 'laterluna'

      assert_nil store.get_player_score(name1)
      assert_nil store.get_player_score(name2)

      # GET player(score)
      get "/players/#{name1}"
      assert last_response.not_found?
      get "/players/#{name2}"
      assert last_response.not_found?

      # POST player(win)
      7.times { post "/players/#{name1}" }
      8.times { post "/players/#{name2}" }

      # GET league
      get '/league'

      assert last_response.ok?
      assert_equal 'application/json', last_response.content_type
      got = JSON.parse last_response.body
      assert_equal [[name2, 8], [name1, 7]], got

      # DELETE player
      delete "/players/#{name2}"
      assert last_response.ok?

      # GET league (again)
      get '/league'

      assert last_response.ok?
      assert_equal 'application/json', last_response.content_type
      got = JSON.parse last_response.body
      assert_equal [[name1, 7]], got

      # POST player(win) (a new player this time)
      8.times { post "/players/#{name3}" }

      # GET league (yet again)
      get '/league'

      assert last_response.ok?
      assert_equal 'application/json', last_response.content_type
      got = JSON.parse last_response.body
      assert_equal [[name3, 8], [name1, 7]], got
    end
  end

end
