module Db
  class Store

    def get_league = fail NotImplementedError
    def record_win(_) = fail NotImplementedError
    def get_player_score(_) = fail NotImplementedError

  end
end
