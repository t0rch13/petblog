$(document).ready(function() {
    $.ajax({
        url: '/api/posts/all', 
        method: 'GET',
        success: function(data) {
            var posts = data.posts;
            var postsDiv = $('#posts');
            postsDiv.empty();

            posts.forEach(function(post) {
                console.log(post);
                var postCard = '<div class="col mb-3">' + // Added margin-bottom CSS
                    '<div class="card">' +
                    '<div class="card-body">' +
                    '<h5 class="card-title">' + post.title + '</h5>' +
                    '<h6>By: <span class="card-subtitle mb-2 text-muted">' + post.username + '</span></h6>' +
                    '<h6>About: <span class="card-subtitle mb-2 text-muted">' + post.petname + '</span></h6>' +
                    '<p class="card-text">' + post.content + '</p>' +
                    '<a href="/post/' + post.id + '" class="btn btn-primary">Read More</a>' +
                    '</div>' +
                    '</div>' +
                    '</div> <br>';
            
                postsDiv.append(postCard);
            });
        }
    });
});