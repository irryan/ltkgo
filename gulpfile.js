var gulp = require('gulp')
var child = require('child_process')
var util = require('gulp-util')

gulp.task('default', ['build', 'test'])

gulp.task('ltkgo:glide', function() {
    var glide = child.spawnSync('glide', ['up'])
    if (glide.stderr.length) {
        var lines = glide.stderr.toString()
            .split('\n').filter(function(line) {
                return line.length
            })
        for (var l in lines)
            util.log(util.colors.red(
                'Error (glide up): ' + lines[l]
            ))
    }

    return glide
})

gulp.task('ltkgo:build', function() {
    var build = child.spawnSync('go', ['build'])
    if (build.stderr.length) {
        var lines = build.stderr.toString()
            .split('\n').filter(function(line) {
                return line.length
            })
        for (var l in lines)
            util.log(util.colors.red(
                'Error (go install): ' + lines[l]
            ))
    }

    return build
})

gulp.task('ltkgo:ginkgo', function() {
    var ginkgo = child.spawnSync('ginkgo', ['-race'])
    if (ginkgo.stderr.length) {
        var lines = ginkgo.stderr.toString()
            .split('\n').filter(function(line) {
                return line.length
            })
        for (var l in lines)
            util.log(util.colors.red(
                'Error (ginkgo): ' + lines[l]
            ))
    }
    if (ginkgo.stdout.length) {
        var lines = ginkgo.stdout.toString()
            .split('\n').filter(function(line) {
                return line.length
            })
        for (var l in lines)
            util.log(util.colors.red(
                'Output (ginkgo): ' + lines[l]
            ))
    }

    return ginkgo
})

gulp.task('build', [
    'ltkgo:glide',
    'ltkgo:build'
])

gulp.task('test', [
    'build',
    'ltkgo:ginkgo'
])