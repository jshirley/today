var gulp = require('gulp');
var del = require('del');
var rename = require('gulp-rename');
var source = require('vinyl-source-stream');
var browserify = require('browserify');
var sass = require('gulp-sass')
var sourcemaps = require('gulp-sourcemaps');
var minifyCSS = require('gulp-minify-css');

gulp.task('watch', function(){
  gulp.watch(["./public/javascript/**/*.jsx", "!./public/dist/*"], ["browserify"]);
  gulp.watch(["./public/stylesheets/**/*.scss", "!./public/dist/*"], ["sass"]);
});

gulp.task('clean', function(cb) {
  del(['./public/dist'], cb);
});

gulp.task('sass', ['clean'], function() {
  gulp.src('./public/stylesheets/main.scss')
  .pipe(sourcemaps.init())
  .pipe(sass())
  .pipe(sourcemaps.write())
  .pipe(gulp.dest('./public/dist/'))
  .pipe(minifyCSS())
  .pipe(rename('main.min.css'))
  .pipe(gulp.dest('./public/dist'));
});

gulp.task('browserify', function() {
    var bundleStream = browserify('./public/javascript/today.jsx')
      .transform('reactify')
      .bundle()
      .pipe(source('bundle.js'));
      return bundleStream.pipe(gulp.dest('./public/dist'));
});
