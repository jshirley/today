var gulp = require('gulp');
var source = require('vinyl-source-stream');
var browserify = require('browserify');

gulp.task('watch', function(){
    gulp.watch(["./public/**/*.jsx", "!./public/javascript/bundle.js"], ["browserify"]);
});

gulp.task('browserify', function() {
    var bundleStream = browserify('./public/javascript/today.jsx')
      .transform('reactify')
      .bundle()
      .pipe(source('bundle.js'));
      return bundleStream.pipe(gulp.dest('./public/javascript/dist'));
});
