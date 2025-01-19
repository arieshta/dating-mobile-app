const users = [
  {
    _id: {
      oid: '67882c786a318f89046afabe'
    },
    username: 'arieshta',
    password: 'arieshta',
    email: 'arieshta@gmail.com',
    fullname: 'Aries Hidayat',
    gender: 'male',
    picture: 'https://avatars.githubusercontent.com/u/16465521?v=4'
  },
  {
    _id: {
      oid: '67882c786a318f89046afac1'
    },
    username: 'nisa',
    password: 'nisa',
    email: 'nisa@gmail.com',
    fullname: 'Nisa',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/12.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afac4'
    },
    username: 'user3',
    password: 'user3',
    email: 'user3@gmail.com',
    fullname: 'User 3',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/22.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afac7'
    },
    username: 'user4',
    password: 'user4',
    email: 'user4@gmail.com',
    fullname: 'User 4',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/32.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afaca'
    },
    username: 'user5',
    password: 'user5',
    email: 'user5@gmail.com',
    fullname: 'User 5',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/45.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afacd'
    },
    username: 'user6',
    password: 'user6',
    email: 'user6@gmail.com',
    fullname: 'User 6',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/67.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afad0'
    },
    username: 'user7',
    password: 'user7',
    email: 'user7@gmail.com',
    fullname: 'User 7',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/54.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afad3'
    },
    username: 'user8',
    password: 'user8',
    email: 'user8@gmail.com',
    fullname: 'User 8',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/72.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afad6'
    },
    username: 'user9',
    password: 'user9',
    email: 'user9@gmail.com',
    fullname: 'User 9',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/95.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afad9'
    },
    username: 'user10',
    password: 'user10',
    email: 'user10@gmail.com',
    fullname: 'User 10',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/98.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afadc'
    },
    username: 'user11',
    password: 'user11',
    email: 'user11@gmail.com',
    fullname: 'User 11',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/63.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afadf'
    },
    username: 'user12',
    password: 'user12',
    email: 'user12@gmail.com',
    fullname: 'User 12',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/89.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afae2'
    },
    username: 'user13',
    password: 'user13',
    email: 'user13@gmail.com',
    fullname: 'User 13',
    gender: 'male',
    picture: 'https://randomuser.me/api/portraits/men/36.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afae5'
    },
    username: 'user14',
    password: 'user14',
    email: 'user14@gmail.com',
    fullname: 'User 14',
    gender: 'female',
    picture: 'https://randomuser.me/api/portraits/women/75.jpg'
  },
  {
    _id: {
      oid: '67882c786a318f89046afae8'
    },
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
    _id: {
      oid: '67882c786a318f89046afae9'
    },
    user_id: {
      oid: '67882c786a318f89046afabe'
    },
    duration: 30,
    purchased_at: new Date(),
    purchase_number: 1
  },
  {
    _id: {
      oid: '67882c786a318f89046afaeb'
    },
    user_id: {
      oid: '67882c786a318f89046afac1'
    },
    duration: 60,
    purchased_at: new Date(),
    purchase_number: 2
  },
  {
    _id: {
      oid: '67882c786a318f89046afaec'
    },
    user_id: {
      oid: '67882c786a318f89046afac4'
    },
    duration: 6,
    purchased_at: new Date(),
    purchase_number: 2
  }
]

const likes = [
  {
    _id: {
      oid: '67882c786a318f89046afaed'
    },
    user_id: {
      oid: '67882c786a318f89046afabe'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afacd'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afaf0'
    },
    user_id: {
      oid: '67882c786a318f89046afac1'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afacd'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afaf3'
    },
    user_id: {
      oid: '67882c786a318f89046afac4'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afacd'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afaf6'
    },
    user_id: {
      oid: '67882c786a318f89046afac7'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afacd'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afaf9'
    },
    user_id: {
      oid: '67882c786a318f89046afaca'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afacd'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afafc'
    },
    user_id: {
      oid: '67882c786a318f89046afacd'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afad0'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afb01'
    },
    user_id: {
      oid: '67882c786a318f89046afad0'
    },
    like_ids: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      },
      {
        oid: '67882c786a318f89046afaca'
      },
      {
        oid: '67882c786a318f89046afacd'
      }
    ]
  }
]

const feeds = [
  {
    _id: {
      oid: '67882c786a318f89046afb06'
    },
    user_id: {
      oid: '67882c786a318f89046afabe'
    },
    feeds: [
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afb09'
    },
    user_id: {
      oid: '67882c786a318f89046afac1'
    },
    feeds: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac4'
      },
      {
        oid: '67882c786a318f89046afac7'
      }
    ]
  },
  {
    _id: {
      oid: '67882c786a318f89046afb0c'
    },
    user_id: {
      oid: '67882c786a318f89046afac4'
    },
    feeds: [
      {
        oid: '67882c786a318f89046afabe'
      },
      {
        oid: '67882c786a318f89046afac1'
      },
      {
        oid: '67882c786a318f89046afac7'
      }
    ]
  }
];
