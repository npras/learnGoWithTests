require 'sinatra/base'
require './db/in_memory_store.rb'

class PlayerServer < Sinatra::Base

  configure do
    set :store, Db::InMemoryStore.new
  end

  get '/players/:name' do |name|
    score = settings.store.get_player_score name
    halt 404 unless score
    body score.to_s
  end

  get '/league' do
    content_type :json
    settings.store.get_league.to_json
  end

  post '/players/:name' do |name|
    settings.store.record_win name
    status 201
  end

end
