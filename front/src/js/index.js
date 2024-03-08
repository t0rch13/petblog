$(document).ready(function() {
    $.ajax({
        url: '/api/posts/all', 
        method: 'GET',
        success: function(data) {
            var posts = data.posts;
            var postsDiv = $('#posts');
            postsDiv.empty();

            posts.forEach(function(post) {
                var postCard = '<div class="col">' +
                    '<div class="card">' +
                    '<div class="card-body">' +
                    '<h5 class="card-title">' + post.title + '</h5>' +
                    '<p class="card-text">' + post.content + '</p>' +
                    '<a href="#" class="btn btn-primary">Read More</a>' +
                    '</div>' +
                    '</div>' +
                    '</div>';

                postsDiv.append(postCard);
            });
        }
    });
});