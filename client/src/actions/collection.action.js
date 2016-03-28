import CollectionConstant from './../constants/collection.constant.js'

const add = (collection) => {
    return {
        type: CollectionConstant.ADD,
        collection
    }
}

const getAll = (collections) => {
    return {
        type: CollectionConstant.GET_ALL,
        collections
    }
}

export default {
    add,
    getAll
}
