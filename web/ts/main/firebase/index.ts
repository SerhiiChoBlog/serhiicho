import firebase from 'firebase/compat/app'
import 'firebase/compat/auth'
import 'firebase/compat/firestore'

firebase.initializeApp({
    apiKey: 'AIzaSyACoEB8A2Wi999RUgqnLBYEmogK0np-8mo',
    authDomain: 'serhii-cc780.firebaseapp.com',
    databaseURL:
        'https://serhii-cc780-default-rtdb.europe-west1.firebasedatabase.app',
    projectId: 'serhii-cc780',
    storageBucket: 'serhii-cc780.appspot.com',
    messagingSenderId: '929831572941',
    appId: '1:929831572941:web:ba82e0779c8620fa5ba084',
})

export default firebase
