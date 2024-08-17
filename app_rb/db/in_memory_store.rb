module Db
  class InMemoryStore

    attr_accessor :h

    def initialize(h = {})
      @h = h
    end

    def get_player_score name
      h[name]
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
