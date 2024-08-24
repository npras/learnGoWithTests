Dir.glob('./db/*.rb').each { |file| require file }
Dir.glob('./lib/{controllers}/*.rb').each { |file| require file }

run PlayerController

# run with:
# rackup -p 4567
