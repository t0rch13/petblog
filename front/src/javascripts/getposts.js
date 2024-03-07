$(document).ready(function () {
    const posts = $('#posts');

    $.ajax({
        url: 'http://localhost:8000/api/posts/',
        method: 'GET',
        success: function (response) {
            let html='';
            response.forEach(post => {
                html += `
                <tr>
                    <th scope="row">
                        ${post.id}
                    </th>
                    <td>
                        <a href="post.html?post=${post.id}">
                            ${post.title}
                        </a>
                    </td>
                    <td>${post.petname}</td>
                    <td>${post.created_at.toLocaleString()}</td>
                </tr>
                `;
            });

            posts.prepend(html);
        },
        error: function (error) {
            console.log(error);
        }
    });
});