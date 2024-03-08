$(document).ready(function() {
    var id = window.location.pathname.split('/').pop(); // extracts the last part of the URL path
    var token = localStorage.getItem('authToken'); 

    $.ajax({
        url: 'http://localhost:8000/api/posts/' + id, 
        type: 'GET',
        dataType: 'json', 
        headers: {
            'Authorization': 'Bearer ' + token 
        },
        success: function(data) {
            console.log(data);
            post = data.post;
            $('#title').text(post.title);
            $('#username').text(post.username);
            $('#petname').text(post.petname);
            $('#content').text(post.content);

            var currentUsername = localStorage.getItem('username'); 
            var postUsername = $('#username').text();

            if (currentUsername === postUsername) {
                $('#editPostButton').show();
                $('#deletePostButton').show();
            }
        },
        error: function(error) {
            console.log(error);
        }
    });

    $('#editPostButton').click(function() {
        var title = $('#title').text();
        var content = $('#content').text();

        $('#editTitle').val(title);
        $('#editContent').val(content);

        $('#editPostModal').modal('show');
    });

    $('#saveChangesButton').click(function() {
        var title = $('#editTitle').val();
        var content = $('#editContent').val();
    
        console.log(title);
        console.log(content);
    
        $.ajax({
            url: 'http://localhost:8000/api/posts/' + id, 
            type: 'PUT',
            contentType: 'application/json', // Specify the content type
            headers: {
                'Authorization': 'Bearer ' + token 
            },
            data: JSON.stringify({ // Stringify the data
                title: title,
                content: content
            }),
            success: function(data) {
                console.log(data);
                window.location.href = '/post/' + id;
            },
            error: function(error) {
                console.log(error);
            }
        });
    
        $('#editPostModal').modal('hide');
    });

    $('#deletePostButton').click(function() {
        var confirmation = confirm("Are you sure you want to delete this post?");
    
        if (confirmation) {
            $.ajax({
                url: 'http://localhost:8000/api/posts/' + id, 
                type: 'DELETE',
                headers: {
                    'Authorization': 'Bearer ' + token 
                },
                success: function(data) {
                    console.log(data);
                    window.location.href = '/'; // Redirect to the homepage or wherever you want
                },
                error: function(error) {
                    console.log(error);
                }
            });
        }
    });
});