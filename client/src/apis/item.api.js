import Store from './../store.js'
import ItemAction from './../actions/item.action.js'

const addToCollection = (collectionUUID, item) => {
    return $.ajax({
        type: "POST",
        url: '/api/collections/' + collectionUUID + '/items',
        contentType: "application/json",
        data: JSON.stringify(item)
    }).done((item) => {
        Store.dispatch(ItemAction.add(item));
    });
}

const getAllInCollection = (collectionUUID) => {
    return $.ajax({
        type: "GET",
        url: '/api/collections/' + collectionUUID + '/items'
    }).done((items) => {
        Store.dispatch(ItemAction.getAll(items));
    });
}

export default {
    addToCollection,
    getAllInCollection
}
