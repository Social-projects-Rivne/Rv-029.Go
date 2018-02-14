import {store} from "../store/configureStore";
import {setErrorMessage, setNotificationMessage} from "../actions/DefaultPageActions";

export default (content, status = false) => {
    if (status) {
        store.dispatch(setNotificationMessage(content))
    } else {
        store.dispatch(setErrorMessage(content))
    }
}