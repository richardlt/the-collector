var gulp = require('gulp'),
    gulpgo = require('gulp-go'),
    clean = require('gulp-clean'),
    webpack = require('webpack'),
    concat = require('gulp-concat'),
    WebpackDevServer = require('webpack-dev-server');

var devConfigWebpack = require('./dev.config.webpack.js'),
    prodConfigWebpack = require('./prod.config.webpack.js');

var watchFiles = {
    server: ['./main.go', './server/**/*.go', './client/dist/templates/*.html']
};

gulp.task('clean-js', function() {
    return gulp.src('./client/dist/js/*', {
        read: false
    }).pipe(clean());
});

gulp.task('bundle-dependencies', function() {
    gulp.src([
        './bower_components/jquery/dist/jquery.min.js'
    ]).pipe(concat('bundle-dependencies.js')).pipe(gulp.dest('./client/dist/js'));
});

gulp.task("bundle-client", function(doneCallBack) {
    webpack(prodConfigWebpack, function(err, stats) {
        doneCallBack();
    });
});

gulp.task('bundle', ['clean-js', 'bundle-dependencies', 'bundle-client']);

gulp.task('webpack-dev-server', ['clean-js', 'bundle-dependencies'], function(callback) {
    var compiler = webpack(devConfigWebpack);
    new WebpackDevServer(compiler, {
        publicPath: '/js/',
        hot: true,
        quiet: true,
        noInfo: true,
        stats: {
            colors: true
        }
    }).listen(8081, 'localhost');
});

gulp.task('go-run', function() {
    go = gulpgo.run('main.go', ['-m', 'dev', 'start'], {
        cwd: __dirname,
        stdio: 'inherit'
    });
});

gulp.task('watch-server', ['go-run'], function() {
    gulp.watch(watchFiles.server).on('change', function() {
        go.restart();
    });
});

gulp.task('start-dev', ['webpack-dev-server', 'watch-server', 'go-run']);

gulp.task('default', ['start-dev']);
