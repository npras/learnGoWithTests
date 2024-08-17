require 'sinatra/base'
require_relative './db/in_memory_store'

class PlayerServer < Sinatra::Base

  configure do
    set :store, Db::InMemoryStore.new
  end

  get '/players/:name' do |name|
    score = settings.store.get_player_score name
    if score
      [200, score.to_s]
    else
      status 404
    end
  end

  post '/players/:name' do |name|
    settings.store.record_win name
    status 201
  end

end
