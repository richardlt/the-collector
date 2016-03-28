import CollectionConstant from './../constants/collection.constant.js'

const initialState = []

const collectionsReducer = (state, action) => {
    if (typeof state === 'undefined') {
        return initialState
    }
    switch (action.type) {
        case CollectionConstant.ADD:
            return state.concat(action.collection);
        case CollectionConstant.GET_ALL:
            return action.collections;
        default:
            return state;
    }
}

export default collectionsReducer
