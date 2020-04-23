function commentPost(button, post_id) {
    content = $(button).siblings(".post-comment-input").val();
    if (content != "") {
        var form = new FormData();
        form.append("content", content);

        $.ajax({
            url: "/api/comments/add/" + parseInt(post_id),
            type: "POST",
            timeout: 0,
            processData: false,
            mimeType: "multipart/form-data",
            contentType: false,
            dataType: 'json',
            data: form,
            success: function(r) {
                $(button).siblings(".post-comment-input").val("");
            },
            error: function(e) {
                notifyUser('Error!', 'Search Failed !', 'danger', 3000);
            }
        });
    }
}