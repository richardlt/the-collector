import Store from './../store.js'
import CollectionAction from './../actions/collection.action.js'
import jquery from 'jquery'

const add = (collection) => {
    return jquery.ajax({
        type: "POST",
        url: '/api/collections',
        contentType: "application/json",
        data: JSON.stringify(collection)
    }).done((collection) => {
        Store.dispatch(CollectionAction.add(collection));
    });
}

const getAll = () => {
    return jquery.ajax({
        type: "GET",
        url: '/api/collections'
    }).done((collections) => {
        Store.dispatch(CollectionAction.getAll(collections));
    });
}

export default {
    add,
    getAll
}
