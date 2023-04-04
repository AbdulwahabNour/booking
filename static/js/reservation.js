'use sttric'

 

let attention = prompt();

(() => {
 
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')


    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {

        form.addEventListener('submit', event => {

            if (!form.checkValidity()) {
                event.preventDefault()
                event.stopPropagation()
            };

            form.classList.add('was-validated')
            event.preventDefault()

        }, false)
    })
})()

 


function notify(msg, msgType) {
    notie.alert({
        type: msgType,
        text: msg,
    })
}

function notifyModal(title, text, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    })
}



function prompt() {



    let toast = function (e) {

        const {
            msg = "",
            icon = "success",
            position = 'top-end',
        } = e

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
        })


    }
    let success = function (e) {

        const {
            msg = "",
            title = "",
            footer = "",
        } = e

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer
        })
    }

    let error = function (e) {

        const {
            msg = "",
            title = "",
            footer = "",
        } = e

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer
        })
    }

    async function custom(e) {
        const {
            msg = "",
            title = "",
        } = e;
// swal.fire returns oject
//     {"isConfirmed": true,
//     "isDenied": false,
//     "isDismissed": false,
//     "value": } 
// const { value: result } means put the value of "value" that returned from Swal.fire in result 
        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            focusConfirm: false,
            backdrop: false,
            showCancelButton: true,
            willOpen: () => {
                if(e.willOpen !== undefined){
                    e.willOpen()
                }
                
            },
            preConfirm: () => {
                if(e.preConfirm !== undefined){
                    return  e.preConfirm()
                }
                
            },
            didOpen: () => {
                if(e.didOpen !== undefined){
                    e.didOpen()
                }
            }

            
        }) 

        if(result){
       
            if(result.dismiss !== Swal.DismissReason.cancel){
               if (result.value !== ""){
                if(e.callback !== undefined){
                    e.callback(result)
                }
                 
               }
            } 
        } 
        



    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

 