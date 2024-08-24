require "minitest/test_task"

# or more explicitly:

Minitest::TestTask.create(:spec) do |t|
  t.libs << "spec"
  t.libs << "lib"
  t.warning = false
  t.test_globs = ["spec/**/*_spec.rb"]
end

task :default => :spec
