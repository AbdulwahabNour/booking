
function notify(msg, msgType) {

    notie.alert({
        type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,

    })
}
function notifyModal(title, msg, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        icon: icon,
        html: msg,
        time: 5,
        confirmButtonText: confirmButtonText,
    })
}

function isValidDate(dateString) {
    var dateRegex = /^\d{4}-\d{2}-\d{2}$/;
    return dateRegex.test(dateString) && !isNaN(Date.parse(dateString));
}

function redirectAfterDelay(url, delay) {
    setTimeout(function () {
        window.location.href = url;
    }, delay);
}