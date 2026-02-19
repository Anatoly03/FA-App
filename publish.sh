
# Generate client
npm run build
sudo cp -r fa-quiz-client/dist/* /var/www/fa-quiz-client/

# Generate server binary
cd fa-quiz-api
go build -o fa-quiz-api src/main.go
cd ..

# Update client
sudo systemctl reload caddy

# Update server
sudo systemctl reload fa-quiz-api
