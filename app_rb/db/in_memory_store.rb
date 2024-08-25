require './db/store.rb'

module Db
  class InMemoryStore < Store

    attr_accessor :h

    def initialize(h = {})
      @h = h
    end

    def get_player_score(name) = h[name]

    def remove_player(name) = h.delete(name)

    def get_league
      h.map { [_1, _2] }.sort { _2[1] <=> _1[1] }
    end


    def record_win name
      if h.key? name
        h[name] += 1
      else
        h[name] = 1
      end
    end

  end
end
