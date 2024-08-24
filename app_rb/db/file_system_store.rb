require 'json'
require './db/store.rb'

module Db
  class FileSystemStore < Store

    attr_accessor :data
    attr_reader :file

    def initialize file
      file.rewind
      @file = file
      @data = read_file(file) || []
    end

    def get_league
      sorted = data.sort { _2['Wins'] <=> _1['Wins'] }
      sorted.map { [ _1['Name'], _1['Wins'] ] }
    end

    def record_win(name)
      if idx = get_player_idx(name)
        data[idx]['Wins'] += 1
      else
        data << { 'Name' => name, 'Wins' => 1 }
      end
      write_file
    end

    def get_player_score name
      idx = get_player_idx name
      return nil unless idx
      data.dig(idx, 'Wins')
    end

    private def read_file(file) = JSON.load(file)
    private def get_player_idx(name) = data.find_index { _1['Name'] == name }

    private def write_file
      file.rewind
      JSON.dump data, file
    end

  end
end
