$(document).ready(function() {
    var id = window.location.pathname.split('/').pop(); // extracts the last part of the URL path
    var token = localStorage.getItem('authToken'); // 

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
        },
        error: function(error) {
            console.log(error);
        }
    });
});