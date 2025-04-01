➀GiikuCamp20 - PostgreSQL Docker 構築手順 & テーブル追加手順

Docker 起動方法
→DockerDesktop 開いてください
→server ディレクトリに移動

docker-compose down
docker-compose up -d
テーブル追加したいとき

DB_NAME=portgonext
psql -h localhost -p 5436 -U portgonext -d portgonext

→ ここで Create 文で追加
