import {
    HANDLE_DRAWER_TOGGLE,
    HANDLE_ERROR_MESSAGE,
    HANDLE_NOTIFICATION_MESSAGE,
    HANDLE_PAGE_TITLE_CHANGE
} from '../constants/defaultPage'

export const toggleDrawer = (state) => {
    return {
        type: HANDLE_DRAWER_TOGGLE,
        payload: state
    }
}
export const changePageTitle = (title) => {
    return {
        type: HANDLE_PAGE_TITLE_CHANGE,
        payload: title
    }
}
export const setErrorMessage = (content) => {
    return {
        type: HANDLE_ERROR_MESSAGE,
        payload: content
    }
}
export const setNotificationMessage = (content) => {
    return {
        type: HANDLE_NOTIFICATION_MESSAGE,
        payload: content
    }
}