import { createStore, applyMiddleware } from 'redux'
import rootReducer from '../reducers'

export function configureStore(initialState) {
  const store = createStore(rootReducer, initialState,
    window.__REDUX_DEVTOOLS_EXTENSION__ &&
    window.__REDUX_DEVTOOLS_EXTENSION__())
  return store
}

export const store = configureStore()
