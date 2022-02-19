# Book Manager API

- [API document]
- [Deployment to Heroku]
- [Deployment to AWS]

## Usage in local environment

1. `git clone git@github.com:ysd-rysn/Book-Manager-API.git` 
2. `cd Book-Manager-API`
3. `cp .env.sample .env`
4. `.env`ファイル内に環境変数を設定する(必要な変数は`.env.sample`を参照)
5. `docker-compose up -d`
6. http://localhost:8080 あてにリクエストを送る([API document]を参照)


[API document]: ysd-rysn.github.io/book-manager-api/dist/index.html
[Deployment to Heroku]: https://ysd-rysn-book-manager.herokuapp.com/
[Deployment to AWS]: ###