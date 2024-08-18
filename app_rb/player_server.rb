require 'sinatra/base'
require './db/in_memory_store.rb'

class PlayerServer < Sinatra::Base

  configure do
    set :store, Db::InMemoryStore.new
  end

  get '/players/:name' do |name|
    score = settings.store.get_player_score name
    if score
      body score.to_s
    else
      status 404
    end
  end

  get '/league' do
    league_h = settings.store.get_league
    content_type :json
    league_h.to_json
  end

  post '/players/:name' do |name|
    settings.store.record_win name
    status 201
  end

end
