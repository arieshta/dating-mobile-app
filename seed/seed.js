const users = [
  {
    username: 'arieshta',
    password: 'arieshta',
    email: 'arieshta@gmail.com',
    fullname: 'Aries Hidayat',
    gender: 'male',
    picture: 'https://avatars.githubusercontent.com/u/16465521?v=4'
  },
  {
    username: 'nisa',
    password: 'nisa',
    email: 'nisa@gmail.com',
    fullname: 'Nisa',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/12.jpg'
  },
  {
    username: 'user3',
    password: 'user3',
    email: 'user3@gmail.com',
    fullname: 'User 3',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/22.jpg'
  },
  {
    username: 'user4',
    password: 'user4',
    email: 'user4@gmail.com',
    fullname: 'User 4',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/32.jpg'
  },
  {
    username: 'user5',
    password: 'user5',
    email: 'user5@gmail.com',
    fullname: 'User 5',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/45.jpg'
  },
  {
    username: 'user6',
    password: 'user6',
    email: 'user6@gmail.com',
    fullname: 'User 6',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/67.jpg'
  },
  {
    username: 'user7',
    password: 'user7',
    email: 'user7@gmail.com',
    fullname: 'User 7',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/54.jpg'
  },
  {
    username: 'user8',
    password: 'user8',
    email: 'user8@gmail.com',
    fullname: 'User 8',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/72.jpg'
  },
  {
    username: 'user9',
    password: 'user9',
    email: 'user9@gmail.com',
    fullname: 'User 9',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/95.jpg'
  },
  {
    username: 'user10',
    password: 'user10',
    email: 'user10@gmail.com',
    fullname: 'User 10',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/98.jpg'
  },
  {
    username: 'user11',
    password: 'user11',
    email: 'user11@gmail.com',
    fullname: 'User 11',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/63.jpg'
  },
  {
    username: 'user12',
    password: 'user12',
    email: 'user12@gmail.com',
    fullname: 'User 12',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/89.jpg'
  },
  {
    username: 'user13',
    password: 'user13',
    email: 'user13@gmail.com',
    fullname: 'User 13',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/36.jpg'
  },
  {
    username: 'user14',
    password: 'user14',
    email: 'user14@gmail.com',
    fullname: 'User 14',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/75.jpg'
  },
  {
    username: 'user15',
    password: 'user15',
    email: 'user15@gmail.com',
    fullname: 'User 15',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/43.jpg'
  }
]

const premium = [
  {
    user_id: '1',
    duration: 30,
    purchased_at: new Date(),
    purchase_number: 1
  },
  {
    user_id: '2',
    duration: 60,
    purchased_at: new Date(),
    purchase_number: 2
  },
  {
    user_id: '3',
    duration: 6,
    purchased_at: new Date(),
    purchase_number: 2
  }
]

const likes = [
  {
    user_id: '1',
    like_ids: ['2', '3', '4', '5', '6', '7']
  },
  {
    user_id: '2',
    like_ids: ['1', '3', '4', '5', '6', '7']
  },
  {
    user_id: '3',
    like_ids: ['1', '2', '4', '5', '6', '7']
  },
  {
    user_id: '4',
    like_ids: ['1', '2', '3', '5', '6', '7']
  },
  {
    user_id: '5',
    like_ids: ['1', '2', '3', '4', '6', '7']
  },
  {
    user_id: '6',
    like_ids: ['1', '2', '3', '4', '5', '7']
  },
  {
    user_id: '7',
    like_ids: ['1', '2', '3', '4', '5', '6']
  }
]

const feed = [
  {
    user_id: '1',
    feeds: ['2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12']
  },
  {
    user_id: '2',
    feeds: ['1', '3']
  },
  {
    user_id: '3',
    feeds: ['1', '2', '4', '5', '6', '7', '8', '9', '10', '11', '12']
  },
  {
    user_id: '4',
    feeds: ['1', '5', '6', '7', '8', '9', '10', '11', '12']
  },
  {
    user_id: '5',
    feeds: ['4', '3']
  },
  {
    user_id: '6',
    feeds: ['1', '2', '3', '4', '5', '7', '8', '9', '10', '11', '12']
  },
  {
    user_id: '7',
    feeds: ['2', '5']
  }
]

db = db.getSiblingDB('dating_app')
db.createCollection('users')
db.createCollection('premium')
db.createCollection('likes')
db.createCollection('feed')
db.users.insertMany(users)
db.premium.insertMany(premium)
db.likes.insertMany(likes)
db.feed.insertMany(feed)
