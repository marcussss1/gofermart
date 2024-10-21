db = db.getSiblingDB('gophermartLoyalty');

// Создание коллекции users и уникального индекса по полю login
db.createCollection('users');
db.users.createIndex({ "login": 1 }, { unique: true });

// Создание коллекции orders и уникального индекса по полю number
db.createCollection('orders');
db.orders.createIndex({ "number": 1 }, { unique: true });
db.orders.createIndex({ "user_id": 1, "uploaded_at": -1 });

// Создание коллекции balances и уникального индекса по полю user_id
db.createCollection('balances');
db.balances.createIndex({ "user_id": 1 }, { unique: true });

// Создание коллекции withdrawals
db.createCollection('withdrawals');
db.withdrawals.createIndex({ "user_id": 1, "processed_at": -1 });

// Определение схемы для коллекции users
db.runCommand({
  collMod: "users",
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["_id", "login", "password_hash"],
      properties: {
        _id: { bsonType: "string" },
        login: { bsonType: "string" },
        password_hash: { bsonType: "string" }
      }
    }
  }
});

// Определение схемы для коллекции orders
db.runCommand({
  collMod: "orders",
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["_id", "number", "user_id", "status", "uploaded_at"],
      properties: {
        _id: { bsonType: "string" },
        number: { bsonType: "string" },
        user_id: { bsonType: "string" },
        status: { bsonType: "string", enum: ["NEW", "PROCESSING", "INVALID", "PROCESSED"] },
        accrual: { bsonType: ["double", "null"] },
        uploaded_at: { bsonType: "date" }
      }
    }
  }
});

// Определение схемы для коллекции balances
db.runCommand({
  collMod: "balances",
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["user_id", "current", "withdrawn"],
      properties: {
        user_id: { bsonType: "string" },
        current: { bsonType: "double" },
        withdrawn: { bsonType: "double" }
      }
    }
  }
});

// Определение схемы для коллекции withdrawals
db.runCommand({
  collMod: "withdrawals",
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["_id", "user_id", "order", "sum", "processed_at"],
      properties: {
        _id: { bsonType: "string" },
        user_id: { bsonType: "string" },
        order: { bsonType: "string" },
        sum: { bsonType: "double" },
        processed_at: { bsonType: "date" }
      }
    }
  }
});

print('Database initialization completed.');
