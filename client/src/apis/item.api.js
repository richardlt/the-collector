import Store from './../store.js'
import ItemAction from './../actions/item.action.js'

const addToCollection = (item, collectionSlug) => {
    return $.ajax({
        type: "POST",
        url: '/api/items',
        headers: {
            'X-COLLECTION-SLUG': collectionSlug
        },
        contentType: "application/json",
        data: JSON.stringify(item)
    }).done((item) => {
        Store.dispatch(ItemAction.add(item));
    });
}

const getAllInCollection = (collectionSlug) => {
    return $.ajax({
        type: "GET",
        url: '/api/items',
        headers: {
            'X-COLLECTION-SLUG': collectionSlug
        }
    }).done((items) => {
        Store.dispatch(ItemAction.getAll(items));
    });
}

export default {
    addToCollection,
    getAllInCollection
}
