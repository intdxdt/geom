var glob = require("glob");
var path = require("path");
var util = require('util');
var _ = require('lodash');
var filepath = require("filepath");

var exec = require('child_process').exec;

module.exports = function (grunt) {
  var root_dir = path.resolve(__dirname + '/..');
  var test_dirs = [
    './mbr',
    './point',
    './linestring',
  ];
  test_dirs = _.map(test_dirs, function (curdir) {
    return path.resolve(path.join(root_dir, curdir))
  });

  var watch_patterns = ['/**/*.go'];

  grunt.initConfig({
    watch: {
      scripts: {
        files  : _.map(watch_patterns, function (pattern) {
          if (pattern[0] == "!") {
            return "!" + root_dir + pattern.slice(1)
          }
          return root_dir + pattern
        }),
        tasks  : ['build_go_geom'],
        options: {
          debounceDelay: 500
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('build_go_geom', function () {
    var cmd = 'go test -cover';

    for (var i = 0; i < test_dirs.length; i++) {
      process.chdir(test_dirs[i]);
      exec(cmd, exec_callback.bind(null, path.basename(test_dirs[i])));
      process.chdir(__dirname)
    }

    function exec_callback(name, error, stdout, stderr) {
      var line = '..........................' + name + '............................';
      console.log(line);
      stdout = stdout + stderr;
      if (_.contains(stdout, "FAIL")) {
        console.log(stdout)
      }
      else {
        stdout = stdout.split('\n').slice(-4);
        console.log(stdout.join('\n'))
      }
    }

  });

  grunt.registerTask('default', ['watch']);
  grunt.task.run(['build_go_geom', 'watch']);

};