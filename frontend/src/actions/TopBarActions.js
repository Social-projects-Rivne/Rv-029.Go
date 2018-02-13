import {
    HANDLE_DRAWER_TOGGLE,
    HANDLE_PAGE_TITLE_CHANGE
} from '../constants/top_bar'

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
