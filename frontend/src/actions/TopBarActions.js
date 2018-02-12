import {
    HANDLE_DRAWER_TOGGLE
} from '../constants/top_bar'

export const toggleDrawer = (state) => {
    return {
        type: HANDLE_DRAWER_TOGGLE,
        payload: state
    }
}
