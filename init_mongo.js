my_db = db.getSiblingDB('development');

my_db.createCollection('user_notification_settings');

my_db.createUser({
    "user": "admin",
    "pwd": "adminpass123",
    "roles": [{
        "role": "readWrite",
        "db": "admin"
    }]
})
