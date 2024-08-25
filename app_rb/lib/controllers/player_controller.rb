require 'sinatra/base'
require './db/file_system_store.rb'

class PlayerController < Sinatra::Base

  configure do
    set :filename, 'game.db.json'
    set :data_file, File.new(settings.filename, File::RDWR|File::CREAT, 0666)
    set :store, Db::FileSystemStore.new(data_file)
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

  delete '/players/:name' do |name|
    settings.store.remove_player name
    status 200
  end

end
