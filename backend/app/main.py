from flask import Flask, render_template, jsonify, request, make_response
from flask_restful import Api, Resource
from flask_cors import CORS
from random import *
# from isbn import RAKUTENAPI
# from sql import sql_arg
# RA = RAKUTENAPI()
# SQLARG = sql_arg()

# *****************
# 触らない
# *****************
# FlaskとVueの連携
app = Flask(__name__, static_folder='../../frontend/app/dist/static', template_folder='../../frontend/app/dist')
# # 日本語
# app.config['JSON_AS_ASCII'] = False
# # CORS=Ajaxで安全に通信するための規約
# api = Api(app)
# CORS(app)
@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def index(path):
    return render_template('index.html')

# *****************
# APIの実装
# *****************
# 本情報取得
# @app.route('/bookinfo', methods=['POST'])
# def getbookinfo():
#     isbn = request.form['isbn']
#     print(isbn)
#     booktitle, bookauthor, bookpublisher = SQLARG.get_bookinfo(isbn)
#     book_info = {
#         'title': booktitle,
#         'author':bookauthor,
#         'publisher':bookpublisher
#     }
#     return jsonify(book_info)

# # 貸出処理
# @app.route('/rental_register', methods=['POST'])
# def rentbook():
#     message = ['success']
#     id = request.form['user_id']
#     # tempisbn = request.form['tempisbn']
#     isbns = request.form['isbn'].split(',')
#     print("aaaa")
#     print(isbns)
#     # if tempisbn !='':
#     #     isbns.append(tempisbn)
#     message = rent(id, isbns)
#     result_data = {
#         'message': message,
#     }
#     return jsonify(result_data)

# # 返却処理
# @app.route('/return_register', methods=['POST'])
# def returnbook():
#     print('start return register')
#     message = ['返却完了']
#     id = request.form['user_id']
#     isbns = request.form['isbn'].split(',')
#     if len(isbns) != 0:
#         print(returnbooks(id, isbns))
#     result_data = {
#         'message': message,
#     }
#     return jsonify(result_data)

# # 本情報取得
# @app.route('/book_list', methods=['POST'])
# def getallbooks():
#     booklist_temp = SQLARG.get_all_book_data()
#     print('本情報一覧表示')
#     booklist = []
#     for i in booklist_temp:
#         print(i)
#         title = i[0]
#         author = i[1]
#         publisher = i[2]
#         rentaluser_name = SQLARG.get_renter_info(i[3])
#         # print(rentaluser_name)
#         bookinfo = {
#             'title': title,
#             'author':author,
#             'publisher':publisher,
#             'rentaluser_name':rentaluser_name
#         }
#         booklist.append(bookinfo)
#     return jsonify(booklist)

# # ユーザ情報
# @app.route('/user_info', methods=['POST'])
# def userinfo():
#     id = request.form['user_id']
#     print(id)
#     ren_books,his_books = SQLARG.get_userinfo(id)
#     user_info = {
#         'ren_books': ren_books,
#         'his_books': his_books
#     }
#     return jsonify(user_info)

# #ユーザー登録
# @app.route('/user_add', methods=['POST'])
# def useradd():
#     id = request.form['user_id']
#     email = request.form['email']
#     values = [id,email]
#     SQLARG.insert_user_data(values)
#     return 'success'

# @app.route('/upuser', methods=['POST'])
# def upuser():
#     print("aaaa")
#     id = request.form['user_id']
#     name = request.form['name']
#     email = request.form['email']
#     SQLARG.update_user_data(id,name,email)
#     return 'success'

# @app.route('/addbook', methods=['POST'])
# def Addbook():
#     print("aaaaaaaaaaaaaaaaaaaaaaaaa")
#     print('start addbook')
#     isbns = request.form['isbn'].split(',')
#     print(isbns)
#     add_book(isbns)

#     return 'success'
# @app.route('/ra_book', methods=['POST'])
# def RA_book():
#     isbn=request.form['isbn']
#     ra_book_info=RA.get_book_info_by_isbn(isbn)
#     print(ra_book_info)
#     if ra_book_info==None:
#         bookinfo ={
#             'title': 'null',
#             'author': 'null',
#             'publisher': 'null'
#         }
#     else:
#         bookinfo = {
#                 'title': ra_book_info[0],
#                 'author':ra_book_info[2],
#                 'publisher':ra_book_info[4]
#             }


#     return jsonify(bookinfo)
# # *****************
# # 関数
# # *****************
# def rent(id, isbns):
#     for isbn in isbns:
#         print("rent関数"+isbn)
#         res=SQLARG.rent_book(id, isbn)
#     return res

# def returnbooks(id, isbns):
#     for isbn in isbns:
#         SQLARG.return_book(id,isbn)
#     return '返却完了'

# def add_book(isbns):
#     for isbn in isbns:
#         book_info=RA.get_book_info_by_isbn(isbn)
#         print(book_info)
#         if book_info ==None:
#             print(isbn)
#             book_info=[]
#             book_info.append('null')
#             book_info.append(isbn)
#             book_info.append('null')
#             book_info.append('null')
#             book_info.append('null')
#             book_info.append('null')
#             book_info.append('null')
#         SQLARG.insert_book_data(values=book_info)

# if __name__ == '__main__':
#     # デバック用
#     ISBN = [9784297106386,9784297100339,9784297100339,9784297106614,9784797392562,9784797392579,9784800712615,9784873117584,9784492557204,9784862760852,9784478061527,9784478068267]
#     # add_book(ISBN)
#     app.run(host='0.0.0.0')

# *****************
# 実行
# *****************

if __name__ == '__main__':
    app.run()
