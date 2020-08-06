# Commands list

+ psa cloudformation deleteMongoStack --stack $stack --type $type
+ psa rds createSnapshot --stack $stack --typerds $type --snapName $snapName
+ psa ecs restart --stack $stack --type $type --service $service