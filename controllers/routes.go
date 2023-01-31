package controllers

func (server *Server) initializeRoutes() {
	route := server.Router.HandleFunc

	// Home Route
	route("/", server.Home).Methods("GET")

	//Customer routes
	route("/customer", server.CustomerController)

	//LoanLimit routes
	route("/loan_limit", server.LoanLimitController)

	//Transaction routes
	route("/transaction", server.TransactionController)
}
