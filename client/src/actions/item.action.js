import ItemConstant from './../constants/item.constant.js'

const add = (item) => {
    return {
        type: ItemConstant.ADD,
        item
    }
}

const getAll = (items) => {
    return {
        type: ItemConstant.GET_ALL,
        items
    }
}

const removeAll = () => {
    return {
        type: ItemConstant.REMOVE_ALL
    }
}

export default {
    add,
    getAll,
    removeAll
}
