require './db/store.rb'

module Db
  class StubMemoryStore < Store

    attr_accessor :h, :win_calls

    def initialize(h = {})
      @h = h
      @win_calls = []
    end

    def get_league
      h.map { [_1, _2] }.sort { _2[1] <=> _1[1] }
    end

    def record_win(name) = win_calls << name
    def get_player_score(name) = h[name]
  end
end
