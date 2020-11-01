// unit of work
// https://martinfowler.com/eaaCatalog/unitOfWork.html
// 借用這個名字, 想不到其他好名字
//
// 已達成功能:
// 1. 只提供事務交易的功能, 協調交給 application layer
//
// 未達成功能:
// 1. 跟蹤變化
// 2. 維護了一個變更列表
// 3. 處理並發
// ...
package uow
