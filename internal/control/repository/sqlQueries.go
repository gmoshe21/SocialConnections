package repository

const (
	queryNewCommunication = `UPDATE social_connections
							SET number_of_communications = (
								  SELECT number_of_communications
								  FROM social_connections
								  WHERE sender = $1 AND recipient = $2
							) + 1
						   WHERE sender = $1 AND recipient = $2;`
	queryFetchConnections = `SELECT TO_JSON (rows) 
		FROM( SELECT MAX(number_of_communications) AS MaxConnections,
		  			MIN(number_of_communications) AS MinConnections,
		  			AVG(number_of_communications) AS AverageConnections,
					(SELECT JSON_AGG(Connections)
						FROM ( SELECT sender, recipient
								FROM social_connections
							) AS Connections
						) 
						FROM social_connections) AS rows;`
)