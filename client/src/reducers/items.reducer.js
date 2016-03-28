import ItemConstant from './../constants/item.constant.js'

const initialState = []

const itemsReducer = (state, action) => {
    if (typeof state === 'undefined') {
        return initialState
    }
    switch (action.type) {
        case ItemConstant.ADD:
            return state.concat(action.item);
        case ItemConstant.GET_ALL:
            return action.items;
        case ItemConstant.REMOVE_ALL:
            return initialState;
        default:
            return state;
    }
}

export default itemsReducer
