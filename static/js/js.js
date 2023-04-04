
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
        confirmButtonText: confirmButtonText,
    })
}
