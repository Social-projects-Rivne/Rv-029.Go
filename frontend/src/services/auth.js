import axios from 'axios'

export default {
    'isAuthorized': () => {
        return (sessionStorage.getItem('token') !== null)
    },
    'logOut': () => {
        sessionStorage.removeItem('token')

        delete axios.defaults.headers.common['Authorization']
    },
    'logIn': (token) => {
        sessionStorage.setItem('token', token)

        axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
    },
    'injectAuthHeader': () => {
        if (sessionStorage.getItem('token') !== null) {
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + sessionStorage.getItem('token');
            return true
        } else {
            return false
        }
    }
}
