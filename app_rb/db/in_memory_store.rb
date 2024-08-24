require './db/store.rb'

module Db
  class InMemoryStore < Store

    attr_accessor :h

    def initialize(h = {})
      @h = h
    end

    def get_league = h
    def get_player_score(name) = h[name]

    def record_win name
      if h.key? name
        h[name] += 1
      else
        h[name] = 1
      end
    end

  end
end
