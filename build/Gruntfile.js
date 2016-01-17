var glob = require("glob");
var util = require('util');
var _ = require('lodash');
var filepath = require("filepath");

var exec = require('child_process').exec;
require('shelljs/global');

var child;

process.mutex = 0;

function get_build_files(pattern, callback) {
  glob(__dirname + pattern, function (err, files) {
    callback(err, files)
  });
}

module.exports = function (grunt) {
  var root_dir = __dirname + '/../lib';
  var watch_patterns = [
    '/**/*.js', '/**/*.sh', '!/**/dist/*'
  ];

  grunt.initConfig({
    asyncfoo: {},
    watch   : {
      scripts: {
        files  : _.map(watch_patterns, function (pattern) {
            if (pattern[0] == "!") {
              return "!" + root_dir + pattern.slice(1)
            }
            return root_dir + pattern
          }),
        tasks  : ['resson_gis_build'],
        options: {
          debounceDelay: 500
        }
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('resson_gis_build', function () {

    get_build_files('/../lib/**/*.sh', function (err, sh_files) {

      if (err) {
        throw err;
      }

      _.each(sh_files, function (sh) {
        var parts = filepath.create(sh).split();
        parts.pop();
        var name = parts[parts.length - 1];
        child = exec(sh, exec_callback.bind(null, name));
      });
    });


    function exec_callback(name, error, stdout, stderr) {
      var line = '..........................' + name + '............................';
      console.log(line);
      if (error !== null) {
        console.log('exec error: ' + error);
        console.log('stderr: ' + stderr);
      }
      else {
        console.log('stdout: ' + stdout);
      }
    }
  });

  grunt.registerTask('default', ['watch']);
  grunt.task.run(['resson_gis_build', 'watch']);

};